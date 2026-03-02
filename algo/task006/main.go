package main

// У нас есть объект [Продавец ID] -> [Список городов, где он осуществляет услуги]
// Необходимо по запрошенным городам вернуть такой же объект только с продавцами, у которых есть
// населенные пункты, лишнее надо откинуть

/* in
sellers = {
    1: ['Москва', 'Самара', 'Ростов'],
    2: ['Москва', 'Самара', 'Ростов', 'Казань', 'Курган', 'Пенза'],
    3: ['Самара', 'Ростов', 'Курган', 'Пенза'],
    4: ['Москва', 'Казань', 'Тула'],
}

cities = ['Москва', 'Казань', 'Тула']

### out
{
    1: [Москва],
    2: [Москва, Казань],
    4: [Москва, Казань, Тула],
}
*/

func filter(sellers map[int][]string, cities []string) map[int][]string {
	res := make(map[int][]string, len(sellers))
	set := make(map[string]struct{}, len(cities))

	for _, city := range cities {
		set[city] = struct{}{}
	}

	for id, cs := range sellers {
		for _, city := range cs {
			if _, exists := set[city]; exists {
				res[id] = append(res[id], city)
			}
		}
	}

	return res
}
