import React, { useEffect, useState } from 'react';
import '../CSS/ProfilePage.css';
import NavBar from '../component/NavBar';
import Login from './Login.jsx';

export function ProfileStatus(){

const [isLogin,setLoginIn]=useState(null)


useEffect(()=>{
    async function checkLogin(){
        try {
            const response = await fetch("/glow-zaar/my-profile-status", {
                                    credentials: "include"
                                        })

            const data =await response.json()
            setLoginIn(data)
        } catch (error) {
            console.log("Error while fetching Cookie Found details")
        }
    }
    checkLogin()
},[])

if(!isLogin){
    return(
        <h2>Loading login status....</h2>
    );
}
if(!isLogin.isFound){
        return(
            <div>
                <NavBar/>
                <Login/>
            </div>
        );
    }
    return <ProfilePage/>
}


export function ProfilePage()
{

    const[profile,setProfile]=useState(null)
    async function logout()
    {
        try {
            const response =await fetch("/glow-zaar/logout",{
                credentials : 'include'
            })
            console.log("logout successfull")
        } catch (error) {
            console.log("error while logout")
        }
    }
    useEffect(()=>{
        async function getProfileDetail(){
            try {
                const response=await fetch("/glow-zaar/my-profile",{
                    credentials : 'include'
                })
                const data=await response.json()
                console.log("profile - ",data)
                setProfile(data)
            } catch (error) {
                console.log("profile error ")
            }
        }
        getProfileDetail()
    },[])

    if(!profile){
        return(
            <h2>Loading Profile ........ </h2>
        )
    }
    return (
    <>
      <NavBar />
      <div className="profile-wrapper">
        <div className="profile-card">
          <h2 className="profile-title">User Profile</h2>
          <div className="profile-details">
            <p><strong>Name : </strong> {profile.name}</p>
            <p><strong>Age : </strong> {profile.age}</p>
            <p><strong>Address : </strong> {profile.address}</p>
            <p><strong>Email : </strong> {profile.email}</p>
          </div>
          <button className="btn logout-btn" onClick={()=>{logout()}}>Logout</button>
        </div>
      </div>
    </>
  );
};

