import { h } from 'preact';
import { useState, useEffect } from "preact/hooks";
import axios from "axios";

export function App(props: any) {
    const [isLoading, setLoading] = useState<boolean>(true);
    const [serverError, setServerError] = useState<any>();
    useEffect(() => {
        async function healthCheck() {
            try {
                await axios.get("http://localhost:8080/ping");
                setLoading(() => false);
            } catch (e) {
                setLoading(() => false);
                setServerError(() => e);
            }
        }
        healthCheck();
    }, []);

    if (isLoading) return (<><h1>Loading!</h1></>);
    if (serverError) return (<><h1>Error connecting to server</h1></>);
    return (<>
        <h1>Let's Backtest!</h1>
    </>);
}
