document.addEventListener("DOMContentLoaded",() => {
    const loginModalToggleButton = document.getElementById("login-modal-toggle")
    if (!loginModalToggleButton){
        return
    }
    
    const body = document.querySelector("body")
    let modalOpen = false;
    let userEmail = "";

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

        const loginModalHeader = document.createElement("h2")
        loginModalHeader.id = "login-modal-header"
        loginModalHeader.textContent = "Login or Signup with Email"

        const loginModalDescriptionOne = document.createElement("p")
        loginModalDescriptionOne.id = "login-modal-description-one"
        loginModalDescriptionOne.textContent = "Enter your email to receive a One Time Password."

        const loginModalDescriptionTwo = document.createElement("p")
        loginModalDescriptionTwo.id = "login-modal-description-two"
        loginModalDescriptionTwo.textContent = "If you don't have an account associated with this email yet, one will be created for you automatically."

        const submitEmailForm = document.createElement("form")
        submitEmailForm.id = "login-modal-form"
        submitEmailForm.addEventListener("submit", submitEmail)

        const emailInput = document.createElement("input")
        emailInput.type="text"
        emailInput.id="login-modal-email-input"

        const emailInputLabel = document.createElement("label")
        emailInputLabel.htmlFor="login-modal-email-input"
        emailInputLabel.textContent="Enter email:"

        const submitEmailInput = document.createElement("input")
        submitEmailInput.id = "login-modal-submit-email-input"
        submitEmailInput.type = "submit"

        const submitOTPForm = document.createElement("form")
        submitOTPForm.id = "login-modal-submit-otp-form"
        submitOTPForm.addEventListener("submit", submitOTP)

        const OTPInput = document.createElement("input")
        OTPInput.id = "login-modal-otp-input"
        OTPInput.type = "text"

        const OTPInputLabel = document.createElement("label")
        OTPInputLabel.htmlFor = "login-modal-otp-input"
        OTPInputLabel.textContent = "Enter your One Time Password"

        const submitOTPInput = document.createElement("input")
        submitOTPInput.type = "submit"

        submitEmailForm.append(emailInputLabel, emailInput, submitEmailInput)
        submitOTPForm.append(OTPInputLabel, OTPInput, submitOTPInput)
        modalContainer.append(closeModalButton, loginModalHeader, loginModalDescriptionOne, loginModalDescriptionTwo, submitEmailForm,  submitOTPForm)
        body.append(modalContainer)
    }

    function destroyModal(){
        modalOpen = false
        const modalContainer = document.getElementById("login-modal-container")
        modalContainer.remove()
    }

    function submitEmail(e){
        e.preventDefault()
        const emailInput = document.getElementById("login-modal-email-input")
        userEmail = emailInput.value
    }

    function submitOTP(e){
        e.preventDefault()
        console.log(userEmail)
    }
})