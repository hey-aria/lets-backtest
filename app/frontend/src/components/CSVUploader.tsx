// CSVUploader.tsx: Allows user to upload their trade data via a .csv file and 
// sends it to the server to be processed and stored
import { h } from "preact";
import { useState } from "preact/hooks";
import axios from "axios";

export default function CSVUploader(props: any) {
    const [serverResponse, setServerResponse] = useState<Number>();
    const [serverError, setServerError] = useState();
    const [fileError, setFileError] = useState<string>();
    const [csv, setCsv] = useState<Blob>();
    const formData = new FormData();
    if (csv) {
        formData.append('data', csv);
    }

    // TODO: I dont know if its better or feasible to do this on the frontend 
    // but it would be nice to validate the shape of the csv as well.
    function handleChange(e: any) {
        if (e.currentTarget.files) {
            if (e.currentTarget.files[0].type !== 'text/csv') {
                setFileError("Invalid file format, please upload a .csv");
            } else {
                setCsv(e.currentTarget.files[0]);
            }
        }
    }

    function handleSubmit(e: any) {
        e.preventDefault();

        async function sendData() {
            try {
                const res = await axios.post("http://localhost:8080/upload/trades-csv", formData);
                console.log(res)
                setServerResponse(res.status);
            } catch (e: any) {
                setServerError(e);
            }
        }
        sendData();
    }

    return (<>
        <form onSubmit={handleSubmit}>
            <input type="file" accept=".csv" onChange={handleChange} />
            <button type="submit">Upload</button>
        </form>
        {/* status TODO: Much better formatting */}
        <>
            {serverError ? (<h1>Internal error</h1>) : fileError ? (<h1>{fileError}</h1>) : serverResponse === 200 ? (<h1>File uploaded!</h1>) : null}
        </>
    </>);
}
