function buttonBefore(button) {
	button.disabled = true
	button.classList.add("loading")
}

function buttonAfter(button) {
	button.disabled = false
	button.classList.remove("loading")
}

function formBefore(form) {
	let button = form.querySelector("button.submit")
	buttonBefore(button)
}

function formAfter(form) {
	let button = form.querySelector("button.submit")
	buttonAfter(button)
}

function fadeInError() {
	$("#error-message").transition("fade in")
}
