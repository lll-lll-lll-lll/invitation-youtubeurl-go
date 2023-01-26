import React, { useState } from 'react'
import styles from "@/styles/Home.module.css"

export default function Index() {
    const [url, setURL] = useState<string>("")
    const [id, setID] = useState<string>("")
    const [invitationCode, setInvitationCode] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const [showPassword, setShowPassword] = useState(false)
    const [clicked, setClicked] = useState<boolean>(false)
    const [getURLError, setGetURLError] = useState("")

    async function GetURL({ id, password, invitationCode }: { id: string, password: string, invitationCode: string }) {
        try {
            console.log("Get URL メソッド実行前")
            const res = await fetch("/api/geturl", {
                method: "POST",
                body: JSON.stringify({ "code": invitationCode, "id": id, "password": password })
            })
            const data = await res.json()
            console.log(data)
            console.log("Get URL メソッド実行後")
            return data
        } catch (err: any) {
            setGetURLError("ID, Password, Codeいずれかが間違えています");
            console.log(err)
        }

    }

    const toggleShowPassword = () => {
        setShowPassword(!showPassword);
    };
    const onChangeID = (e: any) => {
        e.preventDefault()
        setID(e.target.value)
        setClicked(false)
    }
    const onChangePass = (e: any) => {
        e.preventDefault()
        setPassword(e.target.value)
        setClicked(false)
    }
    const onClickForm = async (e: any) => {
        e.preventDefault()
        setClicked(true)
        const data = await GetURL({ id, password, invitationCode })
        setURL(data.message)
    }
    const onChangeInvitationCode = (e: any) => {
        e.preventDefault()
        setInvitationCode(e.target.value)

    }
    return (
        <>

            <div className={styles.section}>

                <div className={styles.wrapper}>
                    <div className={styles.container}>
                        <h2 id="title" className={styles.title}>
                            Get URL
                        </h2>
                        <form className={styles.id_password_form}>
                            <div>
                                <label className={styles.id_password_wrapper_label}>ID</label>
                                <input name="id" id="id" onChange={onChangeID} className={styles.id_form_input} placeholder="ID" />
                            </div>
                            <div>
                                <label className={styles.id_password_wrapper_label}>Password</label>
                                <input
                                    type={showPassword ? 'text' : 'password'}
                                    name="password" id="password"
                                    value={password}
                                    onChange={onChangePass}
                                    onClick={toggleShowPassword}
                                    placeholder="••••••••"
                                    className={styles.id_form_input}
                                />
                            </div>
                            <div>
                                <label className={styles.id_password_wrapper_label}>Invitation Code</label>
                                <input
                                    type="url"
                                    name="url"
                                    id="url"
                                    onChange={onChangeInvitationCode}
                                    className={styles.id_form_input} placeholder="Input Invitation Code" />
                            </div>
                            <button
                                className={styles.id_password_confirm_button}
                                onClick={onClickForm}>
                                Get YoutubeURL
                            </button>
                        </form>
                        <div className={styles.qrcode}>

                            {clicked && (
                                <a href={url}>{url}</a>
                            )}
                        </div>

                    </div>
                </div>
            </div>
        </>
    )
}
