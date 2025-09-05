import Link from "next/link"

export default function Home() {
  return (
    <>
    <h1>Main Page</h1>
    <br />
    <Link href="/products">Product Page</Link>
    <br />
    <Link href="/albums">Albums Page</Link>
    </>
  )
}
