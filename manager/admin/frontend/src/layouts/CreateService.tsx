import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { TextField } from "@mui/material";

export default function CreateService() {
  const onClickCreate = async () => {
    // const response = await fetch("http://localhost:5500/api/containers");
    // const json = await response.json();
    // //  const body = await reader.text();
    // console.log(json);
    await fetch("http://localhost:5500/api/container", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        service: "customer",
        port: "12090",
        artifact:
          "C:\\Users\\info\\Dev\\modular-synthesis-api\\samples\\app\\output",
      }),
    });
  };
  return (
    <Box>
      <ModuleTitle label="Create Service" />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 3,
          maxWidth: "320px",
        }}
      >
        <TextField label="Service name" variant="outlined" />
        <TextField label="Port number" variant="outlined" />
        <TextField
          label="Root directory for execute files"
          variant="outlined"
        />

        <Box sx={{ marginTop: 3, display: "flex", gap: 2 }}>
          <Button variant="contained">Create</Button>
          <Button variant="outlined" color="secondary">
            Cancel
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
