package main

type AuthenticationManager struct {
	mp  map[string]int
	ttl int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{map[string]int{}, timeToLive}
}

func (m *AuthenticationManager) Generate(tokenId string, currentTime int) {
	m.mp[tokenId] = currentTime
}

func (m *AuthenticationManager) Renew(tokenId string, currentTime int) {
	// 合理的设计, 应该需要定时清理一下(?)
	if v, ok := m.mp[tokenId]; ok && v+m.ttl > currentTime {
		m.mp[tokenId] = currentTime
	}
}

func (m *AuthenticationManager) CountUnexpiredTokens(currentTime int) (ans int) {
	for _, t := range m.mp {
		if t+m.ttl > currentTime {
			ans++
		}
	}
	return
}
