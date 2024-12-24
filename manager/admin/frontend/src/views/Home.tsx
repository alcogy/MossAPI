import Header from "../components/Header";
import { Box } from "@mui/material";
import SearviceList from "../layouts/ServiceList";
import TableList from "../layouts/TableList";
import CreateService from "../layouts/CreateService";
import CreateTable from "../layouts/CreateTable";
import InfrastructureInfo from "../layouts/InfrastructureInfo";
import BlankContent from "../layouts/BlankContent";
import { HashRouter, Route, Routes } from "react-router-dom";
import TableDetail from "../layouts/TableDetail";

export default function HomeView() {
  return (
    <>
      <Header />
      <Box sx={{ display: "flex", gap: 4, padding: 2 }}>
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: 3,
            width: "560px",
          }}
        >
          <InfrastructureInfo />
          <SearviceList />
          <TableList />
        </Box>
        <Box sx={{ padding: "16px 0", flex: 1 }}>
          <HashRouter basename="/">
            <Routes>
              <Route path="/" element={<BlankContent />} />
              <Route path="/service" element={<CreateService />} />
              <Route path="/table" element={<CreateTable />} />
              <Route path="/table/:table" element={<TableDetail />} />
            </Routes>
          </HashRouter>
        </Box>
      </Box>
    </>
  );
}
