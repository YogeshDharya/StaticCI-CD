package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func fetchS3URI(bucket, imageName string) string {
	return fmt.Sprintf("https://%s.s3.amazon.com/%s",bucket,imageName)
} 

func updateImageSrc(bucket, htmlPath string, imageNames[] string) error {
	html, err := os.ReadFile(htmlPath)
	if err != nil {
		return fmt.Errorf("failed Reading HTML file : - %v", err)
	}
	operativeHTML := string(html)
	for _, imageName := range imageNames {
		placeHolder := fmt.Sprintf("{{%s}}", imageName)
		s3URI := fetchS3URI(bucket,imageName)
		operativeHTML = strings.ReplaceAll(operativeHTML, placeHolder, s3URI)
	}
	err = os.WriteFile(htmlPath, []byte(operativeHTML), 0644)
	if err != nil {
		return fmt.Errorf("failed to Write into html file🔺Error : - %v", err)
	}
	return nil;
}

func handler(ctx context.Context)error{
	bucket := ""
	sess, err :=  session.NewSession(&aws.Config{
		Region: aws.String("")//will have to look for my s3 bucket region 
	})
	if err != nil {
		return fmt.Errorf("Failed to Establish AWS Session 🔺 :- %v",err)
	}
	svc := s3.New(sess)
	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return fmt.Errorf("Unable to fetch Bucket Items !🔺: %v",err)
	}
	fmt.Print("Objects Successfully Fetched from s3! 🟢 ")
	for _, item := range result.Contents{
		fmt.Printf("Name: %s, Last Modified: %s, Size: %d \n", *item.Key, item.LastModified,*item.Size)
	}
	err := updateImageSrc(bucket,htmlPath,imageNames)
	if err != nil {
		return fmt.Errorf("Error updating Html <img src= !! > 🔺 :%v", err)
	}	

	return nil ;
}

func main(){
	lambda.Start(handler)
}

