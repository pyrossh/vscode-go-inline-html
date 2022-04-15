package main

import (
	"fmt"
)

type M map[string]interface{}

type Todo struct {
	ID string
}

func Html(s string) string {
	return s
}

func Css(s string) string {
	return s
}

var _ = Css(`
	body {
		background: yellow;
	}

	.content {
		font-size: 24px;
		margin: 12px;
	}
`)

func main() {
	res := Html(`
		{{#Page "gromer example"}}
		<div class="flex flex-col justify-center items-center">
			{{#Header "123"}}
			A new link is here
			{{/Header}}
			<h1>Hello this is a h1</h1>
			<h2>Hello this is a h2</h2>
			<h3 x-data="{ message: 'I ❤️ Alpine' }" x-text="message">I ❤️ Alpine</h3>
			<table class="table">
				<thead>
					<tr>
						<th>ID</th>
						<th>Title</th>
						<th>Author</th>
					</tr>
				</thead>
				<tbody id="new-book" hx-target="closest tr" hx-swap="outerHTML swap:0.5s">
					{{#each todos as |todo|}}
					<tr>
						<td>{{todo.ID}}</td>
						<td>Book1</td>
						<td>Author1</td>
						<td>
							<button class="button is-primary">Edit</button>
						</td>
						<td>
							<button hx-swap="delete" class="button is-danger"
								hx-delete="/api/todos/{{todo.ID}}">Delete</button>
						</td>
					</tr>
					{{/each}}
				</tbody>
			</table>
		</div>
		{{/Page}}
	`)
	fmt.Println(res)
}
