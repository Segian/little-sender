import { useEffect, useState } from "react"
import ReactJsonView from '@microlink/react-json-view'

function ShowDetails({ id }){
    const [endpoint, setEndpoint] = useState(null)

    useEffect(() => {
        fetch(`http://localhost:8080/endpoint/${id}`)
            .then(res => res.json())
            .then(data => setEndpoint(data))
            .catch(err => console.error(err));
    }, [id])

    if (!endpoint) return <div>Loading...</div>

    return (
        <div>
            <h3>Endpoint Details</h3>
            <p><strong>Name:</strong> {endpoint.name}</p>
            <p><strong>URL:</strong> {endpoint.url}</p>
            <p><strong>Headers:</strong> {<ReactJsonView src={endpoint.headers} theme={"apathy"} displayDataTypes={false} iconStyle="square" displayObjectSize={false} /> }</p>
            <p><strong>Body:</strong> {<ReactJsonView src={endpoint.body} theme={"apathy"} displayDataTypes={false} iconStyle="square" displayObjectSize={false}/>}</p>
        </div>
    )
}

export default ShowDetails;
