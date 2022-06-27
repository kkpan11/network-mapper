package main

import (
	"context"
	"fmt"
	"github.com/amit7itz/goset"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/otterize/otternose/sniffer/pkg/client"
	"sync"
	"time"
)

type StateKeeper struct {
	intents map[string]*goset.Set[string]
	lock    sync.Mutex
}

func NewStateKeeper() *StateKeeper {
	return &StateKeeper{
		intents: make(map[string]*goset.Set[string]),
		lock:    sync.Mutex{},
	}
}

func (s *StateKeeper) NewIntent(srcIp string, destDns string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.intents[srcIp]; !ok {
		s.intents[srcIp] = goset.NewSet[string](destDns)
	} else {
		s.intents[srcIp].Add(destDns)
	}
}

func (s *StateKeeper) PublishIntents() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.intents) == 0 {
		return
	}
	s.PrintIntents()
	mapperClient := client.NewMapperClient("http://localhost:8080/query")
	results := make([]client.CaptureResultForSrcIp, 0, len(s.intents))
	for srcIp, destinations := range s.intents {
		results = append(results, client.CaptureResultForSrcIp{SrcIp: srcIp, Destinations: destinations.Items()})
	}
	err := mapperClient.ReportCaptureResults(context.TODO(), client.CaptureResults{Results: results})
	if err != nil {
		panic(err)
	}
	s.intents = make(map[string]*goset.Set[string])
}

func (s *StateKeeper) PrintIntents() {
	for ip, dests := range s.intents {
		fmt.Printf("%s:\n", ip)
		for _, dest := range dests.Items() {
			fmt.Printf("\t%s\n", dest)
		}
	}
}

func main() {
	handle, err := pcap.OpenLive("any", 0, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	err = handle.SetBPFFilter("udp port 53")
	if err != nil {
		panic(err)
	}
	statekeeper := NewStateKeeper()
	go func() {
		for true {
			time.Sleep(10 * time.Second)
			statekeeper.PublishIntents()
		}
	}()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer != nil && ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			dns, _ := dnsLayer.(*layers.DNS)
			if dns.OpCode == layers.DNSOpCodeQuery {
				for _, question := range dns.Questions {
					statekeeper.NewIntent(ip.SrcIP.String(), string(question.Name))
				}
			}
		}
	}
}
