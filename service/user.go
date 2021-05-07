package service

import (
	"context"

	"github.com/yongcycchen/mall-api/model/args"
	"github.com/yongcycchen/mall-api/pkg/code"
)

func GenVerifyCode(ctx context.Context, req *args.GenVerifyCodeArgs) (retCode int, verifyCode string) {
	retCode = code.SUCCESS
}
