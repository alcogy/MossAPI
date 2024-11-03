import Header from "../components/Header";
import { Box } from "@mui/material";
import SearviceList from "../layouts/ServiceList";
import TableList from "../layouts/TableList";
import CreateService from "../layouts/CreateService";
import CreateTable from "../layouts/CreateTable";
import BlankContent from "../layouts/BlankContent";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import TableDetail from "../layouts/TableDetail";

const router = createBrowserRouter([
  {
    path: "/",
    element: <BlankContent />,
  },
  {
    path: "/service/",
    element: <CreateService />,
  },
  {
    path: "/table",
    element: <CreateTable />,
  },
  {
    path: "/table/:table",
    element: <TableDetail />,
  },
]);

export default function HomeView() {
  return (
    <>
      <Header />
      <Box sx={{ display: "flex", gap: 4, padding: 2 }}>
        <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
          <SearviceList />
          <TableList />
        </Box>
        <Box sx={{ padding: "16px 0", flex: 1 }}>
          <RouterProvider router={router} />
        </Box>
      </Box>
    </>
  );
}
