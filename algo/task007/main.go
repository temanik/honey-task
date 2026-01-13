package main

/*
Есть дерево категорий товаров, размещаемых на X. Нужно найти заданную подкатегорию и вывести путь до неё
от корневой категории, если она присутствует в дереве
### in
class Category {
  name: string;
  children: Array;
}
root = {
  name: "root",
  children: [
    {
      name: "Бытовая техника",
      children: [
        {
          name: "Телевизоры",
          children: [
            { name: "ЭЛТ", children: [] },
            { name: "LED", children: [] },
            { name: "OLED", children: [] }
          ]
        },
        {
          name: "Холодильники",
          children: [
            { name: "Двухкамерные", children: [] },
            { name: "Однокамерные", children: [] }
          ]
        },
        {
          name: "Утюги",
          children: []
        }
      ]
    },
    {
      name: "Растения",
      children: [
        { name: "Комнатные", children: [] },
        { name: "Садовые", children: [] }
      ]
    }
  ]
}
search_name = "OLED"
### out
Бытовая техника > Телевизоры > OLED
*/
