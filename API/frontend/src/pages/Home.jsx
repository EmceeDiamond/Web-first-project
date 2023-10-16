import Navbar from "../components/Navbar";
import { useEffect, useState } from "react";


export default function Home() {
    const [data, setData] = useState([])

    const callAPI = async () => {
        const res = await fetch('http://localhost:8000/providers')
        const body = await res.json()
        console.log(body)
        setData(body)
    }

    const handleDelete = (id) => {
        return async () => {
            const res = await fetch(`http://localhost:8000/api/employees/${id}`, {
                method: 'DELETE'
            })
            const body = await res.json()
            console.log(body)
            callAPI()
        }
    }

    useEffect(()=>{
        callAPI()
    }, [])

  return (
    <div>
        <Navbar />
        <div className="container">
            <table className="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                        <th scope="col">INN</th>
                        <th scope="col">Contact Details</th>
                        <th scope="col">RF</th>
                    </tr>
                </thead>
                <tbody>
                    {data!=null && data.map((item, index)=> (
                        <tr key={index}>
                            <th scope="row">{item.provider_id}</th>
                            <td>{item.provider_name}</td>
                            <td>{item.INN}</td>
                            <td>{item.contact_details}</td>
                            <td>{item.RF}</td>
                            <td>
                                <a href={`/update/${item.provider_id}`} className="btn btn-warning me-3">Update</a>
                                <button type="button" className="btn btn-danger" onClick={handleDelete(item.provider_id)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    </div>
  )
}