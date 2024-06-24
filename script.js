const form = document.querySelector("#form");
const nameInput = document.querySelector("#name");
const emailInput = document.querySelector("#email");
const passwordInput = document.querySelector("#password");
const passwordInput2 = document.querySelector("#password2");

form.addEventListener("submit", (event) => {
    event.preventDefault();

    // Verifica se o nome está vazio
    if (nameInput.value.trim() === "") {
        MostrarErro(nameInput, "Por favor, preencha o seu Nome corretamente");
        return;
    } else {
        MostrarSucesso(nameInput);
    }

    // Verifica se o email está preenchido e se é válido
    if (emailInput.value.trim() === "" || !isEmailValid(emailInput.value)) {
        MostrarErro(emailInput, "Por Favor, Preencha o seu E-mail corretamente");
        return;
    } else{
        MostrarSucesso(emailInput)
    }

    // Verifica se a senha está preenchida
    if (!ValidatePassword(passwordInput.value, 8)) {
        MostrarErro(passwordInput, "Senha de no mínimo 8 dígitos");
        return;
    } else{
        MostrarSucesso(passwordInput);
    }

    // Verifica se as duas senhas são iguais
    if (!comparePassword()) {
        MostrarErro(passwordInput2, "Senhas devem ser compatíveis");
        return;
    } else{
        MostrarSucesso(passwordInput2)
    }

    // Envie o formulário se todos os campos estiverem corretamente preenchidos
    form.submit();
});

// Função que valida o email
function isEmailValid(email) {
    const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
}

// Função que valida a senha
function ValidatePassword(password, minDigits) {
    return password.length >= minDigits;
}

// Função que compara as senhas
function comparePassword() {
    return passwordInput.value === passwordInput2.value;
}

// Função que indica erro
function MostrarErro(input, message) {
    input.classList.add('erro');
    input.classList.remove('sucesso');
    alert(message);
}

// Função que indica sucesso
function MostrarSucesso(input) {
    input.classList.remove('erro');
    input.classList.add('sucesso');
}
