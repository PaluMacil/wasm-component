package main

import (
	"fmt"
	"github.com/PaluMacil/wasm-component/example"
	"github.com/PaluMacil/wasm-component/wc"
	"go.uber.org/zap"
	stdlog "log"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		stdlog.Fatalf("creating logger: %s", err)
	}
	defer log.Sync()

	appTemplateFetcher := wc.NewTemplateFetcher(example.Assets)
	appTemplateRegistry, err := wc.NewTemplateRegistry("app", appTemplateFetcher)
	if err != nil {
		log.Fatal("creating app template registry", zap.Error(err))
	}
	fmt.Println(appTemplateRegistry)
}
