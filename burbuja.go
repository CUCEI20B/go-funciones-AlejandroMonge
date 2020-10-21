package main

func Burbuja(s []int64) {

	var auxiliar int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				auxiliar = s[i]
				s[i] = s[j]
				s[j] = auxiliar
			}
		}
	}
}

func main() {

}
