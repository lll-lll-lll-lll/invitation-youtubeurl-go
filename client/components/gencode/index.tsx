import { useState } from 'react';
import styles from "@/styles/Home.module.css";
import { useIDPassURL } from '@/hooks/useIDPassURL';
import React from 'react'

function Index() {
    const { clickRes, urlRes, idRes, passwordRes } = useIDPassURL()
    const { clicked, setClicked } = clickRes
    const { url, onChangeURL } = urlRes
    const { id, onChangeID } = idRes
    const { password, onChangePass } = passwordRes
    const [invitationCode, setInvitationCode] = useState<string>("")
    const [showPassword, setShowPassword] = useState(false)
    const [error, setError] = useState<string>("")

    const toggleShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const onClickForm = async (e: any) => {
        e.preventDefault()
        const data = await GenInvitationCode({ id: id, password: password, url: url })
        if (data.message === "400") {
            setError(data.message)
            setClicked(true)
            return
        }
        setError("")
        setInvitationCode(data.message)
        setClicked(true)
    }

    async function GenInvitationCode({ id, password, url }: { id: string, password: string, url: string }) {
        try {
            const res = await fetch("/api/gencode", {
                method: "POST",
                body: JSON.stringify({ "id": id, "password": password, "youtube_url": url })
            })
            const data = await res.json()
            if (res.status === 400) {
                throw new Error("400")
            }
            return data
        } catch (err: any) {
            console.log(err)
            return err
        }
    }

    return (
        <>
            <div className={styles.section}>
                <div className={styles.wrapper}>
                    <div className={styles.container}>
                        <h2 id="title" className={styles.title}>
                            Generate Invitation Code
                        </h2>
                        <form className={styles.id_password_form}>
                            <div>
                                <label className={styles.id_password_wrapper_label}>ID</label>
                                <input name="id" id="id" onChange={onChangeID} className={styles.id_form_input} placeholder="your want ID" />
                            </div>
                            <div>
                                <label className={styles.id_password_wrapper_label}>New Password</label>
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
                                <label className={styles.id_password_wrapper_label}>Youtube URL</label>
                                <input
                                    type="url"
                                    name="url"
                                    id="url"
                                    onChange={onChangeURL}
                                    className={styles.id_form_input} placeholder="https://www.youtube.com/" />
                            </div>
                            <button
                                className={styles.id_password_confirm_button}
                                onClick={onClickForm}>
                                Generate InviTation Code
                            </button>
                        </form>
                        <div className={styles.qrcode}>
                            {clicked && !error ? (
                                <div>{invitationCode}</div>
                            ) : (
                                <div>
                                    {"招待コードの生成に失敗しました"}
                                </div>)}
                        </div>
                    </div>
                </div>
            </div>
        </>
    )

}

export default Index