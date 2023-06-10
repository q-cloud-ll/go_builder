package redis

// PreFix 项目redis统一前缀
const PreFix = "go_builder:"

// getRedisKey 给key加上前缀
func getRedisKey(key string) string {
	return PreFix + key
}
