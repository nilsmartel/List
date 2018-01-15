# List

a simple implementation of Linked Lists in Go, I did this mainly for practicing

---

What's next?

* Create Array from List
* Iterator for List
* Lists are currently Fifo only... quite boring actually
* Sort function? (because, you know, I can. But this will become a shitty task, because `interface{}`)
* Well, you could filter Lists for certain types... type-filtering a List would be usefull
* Better way of Printing Lists
* Concat the shit out of Lists, OF COURSE

so yeah, that's pretty much it. Linked Lists are a terrible idea anyway.

---

Let's take a look at the code!
Lists are implemented using the `interface{}` type, meaning you can store whatever type you want in them. Yeah... that's pretty awesome actually, if you're into that sort of stuff. Idk, too kinky for me tho

```go
list := list.New(0, 1, 2, 3, 4, 5)
list.Push(6)
list.Push("Seven")
for {
	value := list.Pop()
	if value == nil {
		break
	}

	fmt.Println(value)
}
```

this will neatly print out

```
0
1
2
3
4
5
6
Seven
```

nice.
