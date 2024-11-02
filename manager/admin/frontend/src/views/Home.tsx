import Header from "../components/Header";
import { Box } from "@mui/material";
import SearviceList from "../layouts/ServiceList";
import TableList from "../layouts/TableList";
import CreateService from "../layouts/CreateService";
import CreateTable from "../layouts/CreateTable";
import BlankContent from "../layouts/BlankContent";

export default function HomeView() {
  return (
    <>
      <Header />
      <Box sx={{ display: "flex", gap: "32px", padding: "16px" }}>
        <Box sx={{ display: "flex", flexDirection: "column", gap: "24px" }}>
          <SearviceList />
          <TableList />
        </Box>
        <Box sx={{ padding: "16px 0", flex: 1 }}>
          {/* <BlankContent /> */}
          <CreateService />
          {/* <CreateTable /> */}
        </Box>
      </Box>
    </>
  );
}
