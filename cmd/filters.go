package cmd

import "strings"

// buildBlacklists 将黑名单字符串列表转换为用户黑名单和仓库黑名单
func buildBlacklists(blacklist []string) ([]string, []string) {
	var userBlacklist []string
	var repoBlacklist []string
	for _, b := range blacklist {
		if strings.HasPrefix(b, "user:") {
			userBlacklist = append(userBlacklist, strings.TrimPrefix(b, "user:"))
		} else if strings.HasPrefix(b, "repo:") {
			repoBlacklist = append(repoBlacklist, strings.TrimPrefix(b, "repo:"))
		} else {
			userBlacklist = append(userBlacklist, b)
			repoBlacklist = append(repoBlacklist, b)
		}
	}
	return userBlacklist, repoBlacklist
}

// buildWhitelists 将白名单字符串列表转换为用户白名单和仓库白名单
func buildWhitelists(whitelist []string) ([]string, []string) {
	var userWhitelist []string
	var repoWhitelist []string
	for _, w := range whitelist {
		if strings.HasPrefix(w, "user:") {
			userWhitelist = append(userWhitelist, strings.TrimPrefix(w, "user:"))
		} else if strings.HasPrefix(w, "repo:") {
			repoWhitelist = append(repoWhitelist, strings.TrimPrefix(w, "repo:"))
		} else {
			userWhitelist = append(userWhitelist, w)
			repoWhitelist = append(repoWhitelist, w)
		}
	}
	return userWhitelist, repoWhitelist
}
