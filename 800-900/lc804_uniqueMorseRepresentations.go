package main


var translate = []string{
    ".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--..",
}

func uniqueMorseRepresentations(words []string) int {
    var sb strings.Builder
    var cnt = make(map[string]struct{})
    for _, word := range words {
        sb.Reset()
        for i := range word {
            sb.WriteString(translate[word[i]-'a'])
        }
        cnt[sb.String()] = struct{}{}
    }
    return len(cnt)
}