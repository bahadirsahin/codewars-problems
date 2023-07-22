package set3

import . "strings"

func Q75() (o string) {
	for _, s := range []string{"Baby shark", "Mommy shark", "Daddy shark", "Grandma shark", "Grandpa shark", "Let's go hunt"} {
		o += Repeat(s+","+Repeat(" doo", 6)+"\n", 3) + s + "!\n"
	}

	return o + "Run away,…\n"
}
