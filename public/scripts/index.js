function buttonBefore(button) {
	button.disabled = true
	button.classList.add("loading")
}

function buttonAfter(button) {
	button.disabled = false
	button.classList.remove("loading")
}
