import { useState } from "react";
import Navbar from "../components/Navbar";
import {useNavigate} from 'react-router-dom'


export default function AddEmployee() {
    const navigate = useNavigate()
    const [data, setData] = useState({
        element_name: "",
        price: "",
        deadline: "",
    })

    const handleInput = (e) => {
        const name = e.target.name
        const value = e.target.value
        setData({...data, [name]: value})
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const {element_name, price, deadline} = data
        const res = await fetch("http://localhost:8000/add_products", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({element_name, price, deadline})
        })
        const body = await res.json()
        console.log(body)
        navigate("/")
    }


  return (
    <div>
        <Navbar />
        <div className="container">
            <h1 className="text-center">Add Products</h1>
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label htmlFor="element_name" className="form-lable">Name</label>
                    <input type="text" className="form-control" name="element_name" id="element_name" onChange={handleInput} />
                </div>
                <div className="mb-3">
                    <label htmlFor="price" className="form-lable">Price</label>
                    <input type="text" className="form-control" name="price" id="price" onChange={handleInput} />
                </div>
                <div className="mb-3">
                    <label htmlFor="deadline" className="form-lable">Deadline</label>
                    <input type="text" className="form-control" name="deadline" id="deadline" onChange={handleInput} />
                </div>
                <button type="submit" className="btn btn-primary w-100">Submit</button>
            </form>
        </div>
    </div>
  )
}