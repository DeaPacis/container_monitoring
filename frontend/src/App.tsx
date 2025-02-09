import React from "react";
import PingTable from "./components/PingTable";

const App: React.FC = () => {
    return (
        <div>
            <h1 style={{ textAlign: "center" }}>Container Monitoring</h1>
            <PingTable />
        </div>
    );
};

export default App;
