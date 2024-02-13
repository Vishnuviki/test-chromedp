package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	// "log"
	// "time"
)

func main() {
	fmt.Println("HELLO WORLD!")

	// allocatorContext, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.Flag("headless", false),
	// )...)
	// defer cancel()

	dockerUrl := "wss://localhost:9222"

	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), dockerUrl)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	if err := chromedp.Run(ctx, chromedp.Navigate("https://example.com")); err != nil {
		// Handle error
		fmt.Println("Error:", err)
	}

	fmt.Println("END")

	// var html string

	// if err := chromedp.Run(ctx,
	// 	chromedp.Navigate("https://connectivity-wizard-stubbed-showcase.dev.ce.eu-central-1-aws.npottdc.sky/"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[1]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[1]/input", "default"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[2]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[2]/input", "8080"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[3]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[3]/input", "sky.slack.com"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Submit("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	// chromedp.Sleep(time.Millisecond*200),

	// 	chromedp.WaitVisible("/html/body/section/div/div/div/div/div[2]/div/h3"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.InnerHTML("/html/body/section/div/div/div/div/div[2]/div/h3", &html),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("No egress Policy case:", html)

	// ctx, cancel = chromedp.NewContext(allocatorContext)
	// defer cancel()

	// if err := chromedp.Run(ctx,
	// 	chromedp.Navigate("https://connectivity-wizard-stubbed-showcase.dev.ce.eu-central-1-aws.npottdc.sky/"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[1]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[1]/input", "tenant-insights"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[2]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[2]/input", "8080"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[3]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[3]/input", "abc.googleapis.com"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Submit("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	// chromedp.Sleep(time.Millisecond*200),

	// 	chromedp.WaitVisible("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.InnerHTML("/html/body/section/div/div/div/div/div[2]/form/div[1]/div/h4", &html),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("DNS lookup case:", html)

	// ctx, cancel = chromedp.NewContext(allocatorContext)
	// defer cancel()

	// if err := chromedp.Run(ctx,
	// 	chromedp.Navigate("https://connectivity-wizard-stubbed-showcase.dev.ce.eu-central-1-aws.npottdc.sky/"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[1]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[1]/input", "tenant-insights"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[2]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[2]/input", "8080"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Click("/html/body/section/div/div/div/div/div[2]/form/div[3]/input"),
	// 	chromedp.SendKeys("/html/body/section/div/div/div/div/div[2]/form/div[3]/input", "abc.googleapis.com"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Submit("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	// chromedp.Sleep(time.Millisecond*200),

	// 	chromedp.WaitVisible("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.Submit("/html/body/section/div/div/div/div/div[2]/form"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.WaitVisible("/html/body/section/div/div/div/div/div[2]/div/h3"),

	// 	chromedp.Sleep(time.Millisecond*800),

	// 	chromedp.InnerHTML("/html/body/section/div/div/div/div/div[2]/div/h3", &html),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Policy allow case:", html)

}
