// @ts-ignore
import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
    const [serverTime, setServerTime] = useState('');
    const [clientIP, setClientIP] = useState('');

    useEffect(() => {
        fetchServerTime();
        fetchClientIP();
    }, []);

    // @ts-ignore
    const fetchServerTime = async () => {
        try {
            const response = await fetch('/api/time');
            const data = await response.json();
            setServerTime(data.time);
        } catch (error) {
            console.error('Error fetching server time:', error);
        }
    };

    // @ts-ignore
    const fetchClientIP = async () => {
        try {
            const response = await fetch('/api/ip');
            const data = await response.json();
            setClientIP(data.ip);
        } catch (error) {
            console.error('Error fetching client IP:', error);
        }
    };

    return (
        <div className="App">
            <header className="App-header">
                <p>Client IP: {clientIP}</p>
                <p>Server Time in Poland: {serverTime}</p>
            </header>
        </div>
    );
}

export default App;
