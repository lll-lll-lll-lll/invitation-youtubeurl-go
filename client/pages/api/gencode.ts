import { NextApiResponse } from 'next'
import { NextRequest } from 'next/server'

async function handler(req: NextRequest, res: NextApiResponse) {
    console.log(req.body)
    try {
        const response = await fetch("http://localhost:8080/create_invitation_code_guest", {
            method: "POST",
            body: req.body
        })
        if (res.statusCode != 200) {
            throw new Error("")
        }
        const data = await response.json()
        return res.status(201).json(data)
    } catch (err: any) {
        console.log(err)
        return res.status(400).end()
    }
}

export default handler