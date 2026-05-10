function login(){
    var loginData = {
        email = document.getElementById("email").value,
        password = document.getElementById("pw").value
    }
    fetch("/login",{
        method: "POST",
        body: JSON.stringify(loginData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        if(response.ok){
            window.open("student.html", "_self")
        } else{
            throw new Error(response.statusText)
        }
    }).catch(e => alert(e));
}