import Navbar from "../components/Navbar";
import { useEffect, useState } from "react";


export default function Products() {
    const [data, setData] = useState([])

    const callAPI = async () => {
        const res = await fetch('http://localhost:8000/products')
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
                        <th scope="col">ElementId</th>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                        <th scope="col">Quantity</th>
                        <th scope="col">Price</th>
                        <th scope="col">Amount</th>
                        <th scope="col">Deadline</th>
                    </tr>
                </thead>
                <tbody>
                    {data!=null && data.map((item, index)=> (
                        <tr key={index}>
                            <th scope="row">{item.provider_id}</th>
                            <td>{item.element_id}</td>
                            <td>{item.element_name}</td>
                            <td>{item.quantity}</td>
                            <td>{item.price}</td>
                            <td>{item.amount}</td>
                            <td>{item.deadline}</td>
                            <td>
                                <a href={`/update/${item.element_id_id}`} className="btn btn-warning me-3">Update</a>
                                <button type="button" className="btn btn-danger" onClick={handleDelete(item.element_id)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    </div>
  )
}