import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { TextField } from "@mui/material";
import { useState } from "react";
import { API_SERVICE_CREATE } from "../common/constants";

interface ServiceForm {
  name: string;
  port: string;
  artifact: string;
}

const initServiceForm = {
  name: "",
  port: "",
  artifact: "",
};

export default function CreateService() {
  const [form, setForm] = useState<ServiceForm>(initServiceForm);
  const onClickCreate = async () => {
    await fetch(API_SERVICE_CREATE, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        service: form.name,
        port: form.port,
        artifact: form.artifact,
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
        <TextField
          label="Service name"
          variant="outlined"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
        />
        <TextField
          label="Port number"
          variant="outlined"
          value={form.port}
          onChange={(e) =>
            (Number(e.target.value) || e.target.value === "") &&
            setForm({ ...form, port: e.target.value })
          }
        />
        <TextField
          label="Root directory for execute files"
          variant="outlined"
          value={form.artifact}
          onChange={(e) => setForm({ ...form, artifact: e.target.value })}
        />

        <Box sx={{ marginTop: 3, display: "flex", gap: 2 }}>
          <Button variant="contained" onClick={onClickCreate}>
            Create
          </Button>
          <Button variant="outlined" color="secondary" href="/">
            Cancel
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
