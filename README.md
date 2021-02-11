# WIP

`homepage` is intended to be used as a browser's home page, providing a configurable wallpaper and a basic kanban board.

This project gave me a good excuse to focus on new ways of doing my frontend.  

Originally, I wanted to use go's experimental web assembly, since I've worked almost entirely in go for the last year.  Sadly, it doesn't support async, and its value debugging is a nightmare.  I have branched this experiment to [wasm](https://github.com/nathanielwheeler/homepage/tree/wasm) if you want to take a look.

Now, I am using typescript.  This has been on my learning list for a while now, and this will give me an excellent chance to do so.

All of my previous major javascript projects have used the MVC pattern.  In this project, I want to break away from that.  A year of go usage has taught me the power of composition over inheritance, so I'm going to do my best to write classless typescript.  I'm sure I will do a writeup on the process later.