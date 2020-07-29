<script>
import { Button } from "sveltestrap";
import axios from 'axios';
import { Alert } from "sveltestrap";

let visible = true;
let btncolor = "light";
let succolor = "success";
let failcolor = "danger";
var Email = {sent: false, wrongmail: false , val: ""}
var Code  = {same: true , sent : false}
var np = { udpated : true , wrong : false}


function GetRecoveryCode() {
  var email = document.getElementById("email");
  var bodyFormData = new FormData();
  Email.val = email.value;

  bodyFormData.set("email",email.value);
  axios ({
    method:'post',
    url:'http://localhost:8080/forgot',
    data:bodyFormData,
    responseType:'json',
    headers: { 'Content-Type': 'multipart/form-data' }
  })
    .then((response) => {
      var resp = response.data;
      var r = JSON.stringify(resp);
      var obj = JSON.parse(r);
      Email.sent = obj.Response;
      if (!Email.sent) {
        Email.wrongmail = true;
        visible = true;
      }
    })
  email.value= "";
}


function CheckCode() {
  var codeForm = document.getElementById("code");
  var bodyFormData = new FormData();
  bodyFormData.set("code",codeForm.value);
  axios ({
    method:'post',
    url:'http://localhost:8080/forgot',
    data:bodyFormData,
    responseType:'json',
    headers: { 'Content-Type' : 'multipart/form-data' }
  })
    .then((response) => {
      var resp = response.data; var r = JSON.stringify(resp); var obj = JSON.parse(r);
      Code.same = obj.IsEqual;
      if (!Code.same){
        Code.same = false;
        visible = true;
      } else {
        Code.sent = true;
      }
    })
  codeForm.value = "";
}


function submt(e) {
  e.preventDefault();
  GetRecoveryCode()
}

function submt2(e) {
  e.preventDefault();
  CheckCode()
}

function showP(pwF) {
  if (pwF.type === "password") {
    pwF.type = "text";
  } else {
    pwF.type = "password";
  }
}

function ShowPw1() {
  var passw = document.getElementById("pw1");
  showP(passw);
}

function ShowPw2() {
  var passw2 = document.getElementById("pw2");
  showP(passw2); 
}

function UpdatePw() {
  var np1 = document.getElementById("pw1");
  var np2 = document.getElementById("pw2");
  if (np1.value === np2.value) {
    var bodyFormData = new FormData();
    bodyFormData.set("npassword", np1.value)
    bodyFormData.set("email",Email.val)
    axios ({
      method:'post',
      url:'http://localhost:8080/forgot',
      data:bodyFormData,
      responseType:'json',
      headers: { 'Content-Type' : 'multipart/form-data' }
    })
    .then((response) => {
      var resp = response.data; var r = JSON.stringify(resp); var obj = JSON.parse(r);
      np.updated = obj.Updated;
      if (np.updated) {
        window.location.replace("/login");
      }
    })
  } else {
    np.wrong = true;
  }
  np1.value = "";
  np2.value = "";
}

function exeUpdate(e) {
  e.preventDefault();
  UpdatePw()
}
</script>

<main>
  { #if !np.udpated }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      cannot update your password
    </Alert>
  { /if }  
  { #if np.wrong }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Passwords fields are not the same
    </Alert>
  { /if }
  { #if Email.sent }
    <Alert color={succolor} isOpen={visible} toggle={() => (visible = false)}>
       A code with 6 degits has been sent to your email account!
    </Alert>
  { /if }
  { #if Email.wrongmail }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Check if the email you've entered is correct or Not
    </Alert>
  { /if }
  { #if !Code.same }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Sorry the code you have entered is not correct :(
    </Alert>
  { /if }
  { #if Code.sent && Email.val != "" }
    <div class="updatepw">
      <form method="post" id="myForm" on:submit={exeUpdate}>
      <input id="pw1"  required={true} type="password" placeholder="New password" name="passw">
      <span on:click={ShowPw1}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path  d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg><br>
        </span><br>
      <input id="pw2"  required={true} type="password" placeholder="Retype New password" name="passw2">
      <span on:click={ShowPw2}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path  d="M12.015 7c4.751 0 8.063 3.012 9.504 4.636-1.401 1.837-4.713 5.364-9.504 5.364-4.42 0-7.93-3.536-9.478-5.407 1.493-1.647 4.817-4.593 9.478-4.593zm0-2c-7.569 0-12.015 6.551-12.015 6.551s4.835 7.449 12.015 7.449c7.733 0 11.985-7.449 11.985-7.449s-4.291-6.551-11.985-6.551zm-.015 3c-2.209 0-4 1.792-4 4 0 2.209 1.791 4 4 4s4-1.791 4-4c0-2.208-1.791-4-4-4z"/></svg><br>
        </span><br>
        <Button  name="btn" id="sub" {btncolor}>Submit</Button>
    </form>
    </div>

  { :else }
  <form method="post" id="myForm" on:submit={submt}>
    <label id="label"><b><h3>Forgot Your Password ?</h3></b></label>
    <div id="FORM">
    <input id="email" type="email" placeholder="Put Your Email Here" name="email" required={true}>
    <Button name="btn" id="sub" {btncolor}>Submit</Button>
    </div>
  </form>

  <div id="codecheck">
    <form on:submit={submt2}>
    <input id="code" type="number" placeholder="Put You Recovery Code you received on your email here" required={true}>
    <Button name="btn" id="sub" {btncolor}>Submit</Button>
    </form>
  </div>
  { /if }
</main>

<style>
  #label{
    padding-top: 1.5%;
  }
  #codecheck {
    padding:16px;
    padding-top: 5%;
  }
  #FORM {
    padding:16px;
    padding-top: 5%;
  }
	.sub {
		background-color: #8cd7f4;
		color: black;
		padding: 16px 20px;
		margin: 8px 0;
		border: none;
		cursor: pointer;
		width: 50%;
		opacity: 0.9;
	}
	.sub:hover {
  	opacity:1;
	}

  .updatepw {
    position: relative;
    padding:16px;
    padding-top:5%;
  }

  .title {
    padding-top:2%;
    text-align : center;
  }

  input{
    width: 50%;
    padding: 15px;
    margin: 5px 0 22px 0;
    display: inline-block;
    border: none;
    background: #f1f1f1;
  }


  input[type=password]::-ms-reveal {
    display: none;
  }

  input:focus{
    background-color: #ddd;
    outline: none;
  }

  span{
    position: absolute;
    right: 26%;
    padding-top : 1.5%;
  }
</style>
