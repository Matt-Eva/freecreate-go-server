document.addEventListener("DOMContentLoaded",() => {
    const loginModalToggleButton = document.getElementById("login-modal-toggle")
    const body = document.querySelector("body")
    let modalOpen = false;

    loginModalToggleButton.addEventListener("click", () =>{
        if (!modalOpen){
            modalOpen = true
            constructModal()
        } 
    })

    function constructModal(){
        const modalContainer = document.createElement("div")
        modalContainer.id = "login-modal-container"

        const closeModalButton = document.createElement("button")
        closeModalButton.textContent = "close modal"
        closeModalButton.addEventListener("click", destroyModal)

        const loginForm = document.createElement("form")

        modalContainer.append(closeModalButton)
        body.append(modalContainer)
    }

    function destroyModal(){
        modalOpen = false
        const modalContainer = document.getElementById("login-modal-container")
        modalContainer.remove()
    }

    function login(){

    }
})