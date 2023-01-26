import Link from 'next/link'
import React from 'react'
import styles from "@/styles/Header.module.css"

function index() {
    return (
        <header>
            <nav className={styles.nav}>
                <div className={styles.wrapper}>
                    <div className="flex items-center">
                        <span className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">GITC</span>
                    </div>
                    <Link href={'/geturl'}>
                        Get URL
                    </Link>
                    <Link href={'/'}>
                        Generate Invitation Code
                    </Link>
                </div>
            </nav>
        </header >
    )
}

export default index