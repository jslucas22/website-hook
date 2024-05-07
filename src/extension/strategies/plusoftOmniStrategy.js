class PlusoftOmniStrategy {
    constructor(usernameInput, passwordInput) {
        this.usernameInput = usernameInput;
        this.passwordInput = passwordInput;
    }

    login() {
        this.authenticate();
    }

    authenticate() {
        const userInputField = document.querySelector('[data-ng-model="vm.login.user"]');
        const passwordInputField = document.querySelector('[data-ng-model="vm.login.secret"]');

        if (userInputField && (!userInputField.value || userInputField.value !== this.usernameInput)) {
            userInputField.value = this.usernameInput;
            userInputField.dispatchEvent(new Event('input', { bubbles: true }));
        }

        if (passwordInputField && (!passwordInputField.value || passwordInputField.value !== this.passwordInput)) {
            passwordInputField.value = this.passwordInput;
            passwordInputField.dispatchEvent(new Event('input', { bubbles: true }));
        }
    }
}
