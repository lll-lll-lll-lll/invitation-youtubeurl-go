import { Dispatch, SetStateAction, useState } from "react";

type URLRes = {
    url: string,
    setURL: Dispatch<SetStateAction<string>>,
    onChangeURL: (e: any) => void
}

type IDRes = {
    id: string,
    setID: Dispatch<SetStateAction<string>>,
    onChangeID: (e: any) => void
}

type PassRes = {
    password: string,
    setPassword: Dispatch<SetStateAction<string>>,
    onChangePass: (e: any) => void
}


export function useURL(setClicked: (value: SetStateAction<boolean>) => void): URLRes {
    const [url, setURL] = useState<string>("")
    const onChangeURL = (e: any) => {
        e.preventDefault()
        setURL(e.target.value)
        setClicked((state: boolean) => !state)
    }
    return { url, setURL, onChangeURL }
}

export function useID(setClicked: (value: SetStateAction<boolean>) => void): IDRes {
    const [id, setID] = useState<string>("")
    const onChangeID = (e: any) => {
        e.preventDefault()
        setID(e.target.value)
        setClicked((state: boolean) => !state)
    }
    return { id, setID, onChangeID }
}

export function usePassWord(setClicked: (value: SetStateAction<boolean>) => void): PassRes {
    const [password, setPassword] = useState<string>("")
    const onChangePass = (e: any) => {
        e.preventDefault()
        setPassword(e.target.value)
        setClicked(false)
    }
    return { password, setPassword, onChangePass }
}

type ALLRea = {
    clickRes: { clicked: boolean, setClicked: Dispatch<SetStateAction<boolean>> },
    urlRes: URLRes,
    idRes: IDRes,
    passwordRes: PassRes
}

export function useIDPassURL(): ALLRea {
    const [clicked, setClicked] = useState<boolean>(false)
    const { url, setURL, onChangeURL } = useURL(setClicked)
    const { id, setID, onChangeID } = useID(setClicked)
    const { password, setPassword, onChangePass } = usePassWord(setClicked)
    return {
        clickRes: { clicked: clicked, setClicked: setClicked },
        urlRes: { url: url, setURL: setURL, onChangeURL: onChangeURL },
        idRes: { id, setID: setID, onChangeID: onChangeID },
        passwordRes: { password: password, setPassword: setPassword, onChangePass: onChangePass }
    }
}