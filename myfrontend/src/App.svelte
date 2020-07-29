<script>
  import Home from "./component/Home.svelte";
  import Login from "./component/Login.svelte";
  import Register from "./component/Register.svelte";
  import Error from "./component/Error.svelte";
  import Forgotpw from "./component/Forgotpw.svelte";
  import Credentials from "./component/Credentials.svelte";
  import CredFile from "./component/CredFile.svelte";
  import GenPassw from "./component/GenPassw.svelte";
  import { Button } from "sveltestrap";
  import { Alert } from "sveltestrap";
  import router from "page";
  import Cookies from "js-cookie";
  import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink,
    UncontrolledDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem
  } from 'sveltestrap';
  import axios from 'axios';

  const color = "info";

  let visible = true;
  let isOpen = false;
  let btncolor = "light";
  let loggedout = true;
  var timeleft = 20; 
  let url = 'http://localhost:8080/logout'
  let timeout;
  let decrease;
  let pop = false ;
  const cookiename = "session";

  function handleUpdate(event) {
    isOpen = event.detail.isOpen;
  }
  
  let page;

  router("/", function(){
    page=Home
  }); 
  router("/login", () => page=Login);
  router("/register", () => page=Register); 
  router("/forgotpw", () => page=Forgotpw);
  router("/credentials",() => page=Credentials);
  router("/credfile",() => page=CredFile);
  router("/generatepw",() => page=GenPassw);
  router("/*",() => page=Error); 
  router.start();
  

function logout() {
  axios({
    method:'get',
    url:url,
    responseType:'json',
    withCredentials:true
  })
    .then(function(response){
      var resp = response.data;
      var r = JSON.stringify(resp);
      var obj = JSON.parse(r);
      loggedout = obj.IsLogout;
    });
}
//TODO: add a countdown box
  function setup() {
    if (pop === false) {
      window.addEventListener("mousemove", resetTimer, false);
      window.addEventListener("mousedown", resetTimer, false);
      window.addEventListener("keypress", resetTimer, false);
      window.addEventListener("DOMMouseScroll", resetTimer, false);
      window.addEventListener("mousewheel", resetTimer, false);
      window.addEventListener("touchmove", resetTimer, false);
      window.addEventListener("MSPointerMove", resetTimer, false);
      tracktime();
    } 
  }
  setup();

  function RemoveEvents() {
    if (pop) {
      window.removeEventListener("mousemove", resetTimer, false);
      window.removeEventListener("mousedown", resetTimer, false);
      window.removeEventListener("keypress", resetTimer, false);
      window.removeEventListener("DOMMouseScroll", resetTimer, false);
      window.removeEventListener("mousewheel", resetTimer, false);
      window.removeEventListener("touchmove", resetTimer, false);
      window.removeEventListener("MSPointerMove", resetTimer, false);
    }
  }

  function tracktime() {
    if (Cookies.get(cookiename) != null) {
      timeout = window.setTimeout(goInactive,600000)
    }
  }
  
  function resetTimer(){
    window.clearTimeout(timeout);
    window.clearInterval(decrease);
    goActive();
  }

  function goInactive() {
    PopOut();
    RemoveEvents();
  }

  function goActive() {
    tracktime()
  }

  function Runtimer() {
    if (pop) {
      decrease = window.setInterval(function() {
        if (timeleft != 0) {
          timeleft-- ;
        } else {
          logout();
          window.location.replace("/")
        }
      },1000)
    }
  }
  


  function PopOut() {
    var modal = document.getElementById("myModal");
    // Get the <span> element that closes the modal
    var span = document.getElementsByClassName("close")[0];
    // When the user clicks on the button, open the modal
    modal.style.display = "block";
    pop = true;
    Runtimer();

    var cancelB = document.getElementById("cancel");
    cancelB.onclick = function() {
      modal.style.display = "none";
      timeleft = 10;
      pop = false ;
      setup();
    }

    var okB = document.getElementById("okL");
    okB.onclick = function() {
      logout();
      window.location.replace("/")
    }
    
    // When the user clicks on <span> (x), close the modal
    span.onclick = function() {
      modal.style.display = "none";
      timeleft = 10;
      pop = false ;
      setup();
    }
    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
      if (event.target == modal) {
        modal.style.display = "none";
        timeleft = 10;
        pop = false;
        setup();
      }
    }

  }

</script>

<main>
  <div id="nav">
  <Navbar color="dark Navbar" light expand="md">
    <NavbarBrand href="/">🦂PWM</NavbarBrand>
    <NavbarToggler on:click={() => (isOpen = !isOpen)} />
    <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
      <Nav class="ml-auto" navbar>
        <NavItem>
          <NavLink href="/login">login</NavLink>
        </NavItem>
        <NavItem>
          <NavLink href="/register">register</NavLink>
        </NavItem>
        <UncontrolledDropdown nav inNavbar>
          <DropdownToggle nav caret>Account</DropdownToggle>
          <DropdownMenu right>
             <DropdownItem on:click={logout}>logout</DropdownItem>
          </DropdownMenu>
        </UncontrolledDropdown>
      </Nav>
    </Collapse>
  </Navbar>
  </div>
  { #if !loggedout }
  <Alert color="success" isOpen={visible} toggle={() => (visible = false)}>
    logged out succesfully
  </Alert>
  { /if }
  <!-- The Modal -->
  <div id="myModal" class="modal">
    <!-- Modal content -->
    <div class="modal-content">
      <span class="close">&times;</span>
      <h2>You have been idle for a quiet amount of time!</h2>
      <h3>you are gonna be logged out in {timeleft}</h3>
      <div class="btns">
        <Button name="btn" id="cancel" {btncolor}>Cancel</Button>
        <Button name="btn" id="okL" {btncolor}>Ok</Button>
      </div> 
    </div>

  </div>
<svelte:component this={page} />
</main>

<style>
  .router-link-active {
    color: blue;
  }
	main {
		text-align: center;
		/*padding: 1em;*/
    padding-bottom:15%;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

  .modal {
    display: none; /* Hidden by default */
    position: fixed; /* Stay in place */
    z-index: 1; /* Sit on top */
    left: 0;
    top: 0;
    width: 100%; /* Full width */
    height: 100%; /* Full height */
    overflow: auto; /* Enable scroll if needed */
    background-color: rgb(0,0,0); /* Fallback color */
    background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
  }

  /* Modal Content/Box */
  .modal-content {
    background-color: #fefefe;
    margin: 15% auto; /* 15% from the top and centered */
    padding: 20px;
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
  }
  .close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
  }

  .close:hover,
  .close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
  }

</style>
