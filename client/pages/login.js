import { useEffect,useState } from 'react'
import { LoginForm } from "../components/LoginForm"
import Link from 'next/link'
import Router from 'next/router'

export default function Login() {
  useEffect(()=>{
    // fetchFormData();
  },[])
  return (
    <div className="container">
      <main>
        <LoginForm/>
        <Link href="/">
          <a>トップページへ戻る</a>
        </Link>
      </main>

      <style jsx>{`
        .container {
          min-height: 100vh;
          padding: 0 0.5rem;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
        }
      `}</style>

      <style jsx global>{`
        html,
        body {
          padding: 0;
          margin: 0;
          font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
            Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
            sans-serif;
        }

        * {
          box-sizing: border-box;
        }
      `}</style>
    </div>
  )
}
