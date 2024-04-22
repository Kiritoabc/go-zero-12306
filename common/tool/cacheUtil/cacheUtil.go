package cacheUtil

import (
	"fmt"
	"strings"
)

const SPLICING_OPERATOR string = "_"

/**
 * 构建缓存标识
 *
 * @param keys
 * @return
 */

func BuildKey(keys []string) (string, error) {
	if len(keys) == 0 {
		return "", fmt.Errorf("keys is empty")
	}
	return strings.Join(keys, SPLICING_OPERATOR), nil
}
