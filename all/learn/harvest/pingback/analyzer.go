package pingback

type Analyzer interface {
	Analyze(pb *Pingback)
}

// SimpleAnalyzer 简易实现
type SimpleAnalyzer struct {
	Testers     []Filter
	NewConsumer func() Consumer

	consumers map[string]Consumer
}

func (analyzer *SimpleAnalyzer) Analyze(pb *Pingback) {
	if !analyzer.test(pb) {
		return
	}
	consumer, ok := analyzer.consumers[pb.SN]
	if !ok {
		consumer = analyzer.NewConsumer()
	}
	consumer.Consume(pb)
	analyzer.consumers[pb.SN] = consumer
}

// Consumer pingback消费
type Consumer interface {
	Consume(pb *Pingback)
}

// Filter pingback过滤
type Filter interface {
	Filter(pb *Pingback) bool
}

func (analyzer *SimpleAnalyzer) test(pb *Pingback) bool {
	for _, tester := range analyzer.Testers {
		if !tester.Filter(pb) {
			return false
		}
	}
	return true
}
