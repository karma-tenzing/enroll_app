function signUp(){
    // get form data   
    var signUpData = {
        firstname : document.getElementById("fname").value,
        lastname : document.getElementById("lname").value,
        email : document.getElementById("email").value,
        password : document.getElementById("pw1").value,
        pw : document.getElementById("pw2").value
    }
    if(signUpData.password !== signUpData.pw){
        alert("Password doesnot match ")
        return
    }
    fetch("/signup", {
        method: "POST",
        body: JSON.stringify(signUpData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response =>{
        if(response.status == 201){
            window.open("index.html", "_self")
        } else{
            throw new Error(response.statusText)
        }
    }).catch(e => alert(e));
}