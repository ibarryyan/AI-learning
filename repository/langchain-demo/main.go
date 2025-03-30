package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
)

func main() {
	//Simple()
	Input()
	//ReasoningPlus()
}

func Simple() {
	// Initialize the OpenAI client with Deepseek model
	llm, err := openai.New(
		openai.WithBaseURL("https://api.deepseek.com"),
		openai.WithModel("deepseek-chat"),
		openai.WithToken(""), // 填写自己的API Key
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Create messages for the chat
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你现在模仿曹操，以曹操的口吻和风格回答问题，要展现出曹操的霸气与谋略"),
		llms.TextParts(llms.ChatMessageTypeHuman, "赤壁之战打输了怎么办？"),
	}

	// Generate content with streaming to see both reasoning and final answer in real-time
	fmt.Print("曹孟德：")
	completion, err := llm.GenerateContent(
		ctx,
		content,
		llms.WithMaxTokens(2000),
		llms.WithTemperature(0.7),
		llms.WithStreamingReasoningFunc(func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {
			contentColor := color.New(color.FgCyan).Add(color.Bold)
			if len(chunk) > 0 {
				_, err := contentColor.Printf("%s", string(chunk))
				if err != nil {
					return err
				}
			}
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(completion.Choices) > 0 {
		choice := completion.Choices[0]
		fmt.Printf("\nFinal Answer:\n%s\n", choice.Content)
	}
}

func Input() {
	// Initialize the OpenAI client with Deepseek model
	llm, err := openai.New(
		openai.WithBaseURL("https://api.deepseek.com"),
		openai.WithModel("deepseek-chat"),
		openai.WithToken(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 初始系统消息
	systemMessage := llms.TextParts(llms.ChatMessageTypeSystem, "你现在模仿曹操，以曹操的口吻和风格回答问题，要展现出曹操的霸气与谋略。")
	content := []llms.MessageContent{systemMessage}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("闫同学：")
		scanner.Scan()
		question := scanner.Text()

		if question == "exit" {
			break
		}

		// 添加新的用户问题
		userMessage := llms.TextParts(llms.ChatMessageTypeHuman, question)
		content = append(content, userMessage)

		fmt.Print("曹孟德：")
		// Generate content with streaming to see both reasoning and final answer in real-time
		completion, err := llm.GenerateContent(
			ctx,
			content,
			llms.WithMaxTokens(2000),
			llms.WithTemperature(0.7),
			llms.WithStreamingReasoningFunc(func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {

				contentColor := color.New(color.FgCyan).Add(color.Bold)

				if len(chunk) > 0 {
					_, err := contentColor.Printf("%s", string(chunk))
					if err != nil {
						return err
					}
				}
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		// 将回复添加到历史消息中
		if len(completion.Choices) > 0 {
			choice := completion.Choices[0]
			assistantMessage := llms.TextParts(llms.ChatMessageTypeHuman, choice.Content)
			content = append(content, assistantMessage)
		}
	}
}

func Reasoning() {
	// Initialize the OpenAI client with Deepseek model
	llm, err := openai.New(
		openai.WithBaseURL("https://api.deepseek.com"),
		openai.WithModel("deepseek-reasoner"),
		openai.WithToken(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 初始系统消息
	systemMessage := llms.TextParts(llms.ChatMessageTypeSystem, "你现在模仿曹操，以曹操的口吻和风格回答问题，要展现出曹操的霸气与谋略。")
	content := []llms.MessageContent{systemMessage}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("闫同学：")
		scanner.Scan()
		question := scanner.Text()

		if question == "exit" {
			break
		}

		// 添加新的用户问题
		userMessage := llms.TextParts(llms.ChatMessageTypeHuman, question)
		content = append(content, userMessage)

		isPrint := false
		fmt.Print("曹孟德：")
		// Generate content with streaming to see both reasoning and final answer in real-time
		completion, err := llm.GenerateContent(
			ctx,
			content,
			llms.WithMaxTokens(2000),
			llms.WithTemperature(0.7),
			llms.WithStreamingReasoningFunc(func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {
				contentColor := color.New(color.FgCyan).Add(color.Bold)
				reasoningColor := color.New(color.FgYellow).Add(color.Bold)

				if !isPrint {
					isPrint = true
					fmt.Print("[思考中]")
				}

				if len(reasoningChunk) > 0 {
					_, err := reasoningColor.Printf("%s", string(reasoningChunk))
					if err != nil {
						return err
					}
				}

				if len(chunk) > 0 {
					_, err := contentColor.Printf("%s", string(chunk))
					if err != nil {
						return err
					}
				}
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		// 将回复添加到历史消息中
		if len(completion.Choices) > 0 {
			choice := completion.Choices[0]
			assistantMessage := llms.TextParts(llms.ChatMessageTypeHuman, choice.Content)
			content = append(content, assistantMessage)
		}
	}
}

func ReasoningPlus() {
	// Initialize the OpenAI client with Deepseek model
	llm, err := openai.New(
		openai.WithBaseURL("https://api.deepseek.com"),
		openai.WithModel("deepseek-reasoner"),
		openai.WithToken(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 初始系统消息
	systemMessage := llms.TextParts(llms.ChatMessageTypeSystem, "你现在模仿曹操，以曹操的口吻和风格回答问题，要展现出曹操的霸气与谋略。")
	content := []llms.MessageContent{systemMessage}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("闫同学：")
		scanner.Scan()
		question := scanner.Text()

		if question == "exit" {
			break
		}

		// 添加新的用户问题
		userMessage := llms.TextParts(llms.ChatMessageTypeHuman, question)
		content = append(content, userMessage)

		res := make(chan string)
		reasoning := make(chan string)

		fmt.Print("曹孟德：")
		go func() {
			i, j := 0, 0
			fmt.Print(i, j)
			for {
				select {
				case r := <-res:
					if i == 0 {
						i++
						fmt.Print("【思考中】")
					}
					fmt.Print(r)
				case a := <-reasoning:
					if j == 0 {
						j++
						fmt.Print("【回答中】")
					}
					fmt.Print(a)
				}
			}
		}()

		// Generate content with streaming to see both reasoning and final answer in real-time
		completion, err := llm.GenerateContent(
			ctx,
			content,
			llms.WithMaxTokens(2000),
			llms.WithTemperature(0.7),
			llms.WithStreamingReasoningFunc(func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {
				//contentColor := color.New(color.FgCyan).Add(color.Bold)
				//reasoningColor := color.New(color.FgYellow).Add(color.Bold)

				if len(reasoningChunk) > 0 {
					reasoning <- string(reasoningChunk)
					//_, err := reasoningColor.Printf("%s", string(reasoningChunk))
					//if err != nil {
					//    return err
					//}
				}

				if len(chunk) > 0 {
					res <- string(chunk)
					//_, err := contentColor.Printf("%s", string(chunk))
					//if err != nil {
					//    return err
					//}
				}
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}

		// 将回复添加到历史消息中
		if len(completion.Choices) > 0 {
			choice := completion.Choices[0]
			assistantMessage := llms.TextParts(llms.ChatMessageTypeHuman, choice.Content)
			content = append(content, assistantMessage)
		}
	}
}
