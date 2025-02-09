import React, { useEffect, useState } from "react";

interface PingResult {
    ip_address: string;
    response_time: number;
    last_checked: string;
}

const backendUrl = process.env.REACT_APP_API_URL || "http://localhost:8090/ping-table";

const PingTable: React.FC = () => {
    const [pingResults, setPingResults] = useState<PingResult[]>([]);

    const fetchPingResults = async () => {
        try {
            const response = await fetch(backendUrl + "/ping-table");
            if (!response.ok) throw new Error("Failed to fetch data");

            const data = await response.json();

            const filteredData = data.map((item: any) => ({
                ip_address: item.ip_address,
                response_time: item.response_time,
                last_checked: item.last_checked
            }));

            setPingResults(filteredData);
        } catch (error) {
            console.error("Error fetching ping results:", error);
        }
    };

    useEffect(() => {
        fetchPingResults();
        const interval = setInterval(fetchPingResults, 10000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div>
            <table style={{ borderCollapse: "collapse", width: "100%" }}>
                <thead>
                <tr>
                    <th style={tableHeaderStyle}>IP Address</th>
                    <th style={tableHeaderStyle}>Response Time (us)</th>
                    <th style={tableHeaderStyle}>Last Checked</th>
                </tr>
                </thead>
                <tbody>
                {pingResults.map((result, index) => (
                    <tr key={index}>
                        <td style={tableCellStyle}>{result.ip_address}</td>
                        <td style={tableCellStyle}>{result.response_time}</td>
                        <td style={tableCellStyle}>{result.last_checked}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

const tableHeaderStyle: React.CSSProperties = {
    border: "1px solid black",
    padding: "8px",
    backgroundColor: "#0075FD",
    color: "white",
    textAlign: "left"
};

const tableCellStyle: React.CSSProperties = {
    border: "1px solid black",
    padding: "8px"
};

export default PingTable;
