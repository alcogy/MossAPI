import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { TextField } from "@mui/material";
import { useState } from "react";
import { API_SERVICE_CREATE } from "../common/constants";

interface ServiceForm {
  name: string;
  artifact: string;
  options: string;
  command: string;
}

const initServiceForm = {
  name: "",
  artifact: "",
  options: "",
  command: "",
};

export default function CreateService() {
  const [form, setForm] = useState<ServiceForm>(initServiceForm);
  const onClickCreate = async () => {
    const result = await fetch(API_SERVICE_CREATE, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        service: form.name,
        artifact: form.artifact,
        options: form.options,
        command: form.command,
      }),
    });

    console.log(result);
    setForm(initServiceForm);
  };

  return (
    <Box>
      <ModuleTitle label="Create Service" />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 3,
          maxWidth: "480px",
        }}
      >
        <TextField
          label="Service name"
          variant="outlined"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
        />
        <TextField
          label="Root directory for artifact"
          variant="outlined"
          value={form.artifact}
          onChange={(e) => setForm({ ...form, artifact: e.target.value })}
        />
        <TextField
          label="Optional Dockerfile Commands."
          variant="outlined"
          multiline
          rows={6}
          value={form.options}
          onChange={(e) => setForm({ ...form, options: e.target.value })}
        />
        <TextField
          label="Command when start contaier."
          variant="outlined"
          value={form.command}
          onChange={(e) => setForm({ ...form, command: e.target.value })}
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
