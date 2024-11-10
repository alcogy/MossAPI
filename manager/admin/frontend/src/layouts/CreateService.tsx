import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { TextField } from "@mui/material";
import { useState } from "react";
import { API_GET_SERVICES, API_SERVICE_CREATE } from "../common/constants";
import { loadingState, serviceListState } from "../state/atoms";
import { useRecoilState } from "recoil";
import { Service } from "../state/models";

interface ServiceForm {
  name: string;
  artifact: string;
  options: string;
  execute: string;
}

const initServiceForm = {
  name: "",
  artifact: "",
  options: "",
  execute: "",
};

export default function CreateService() {
  const [form, setForm] = useState<ServiceForm>(initServiceForm);
  const [_serviceList, setServiceList] = useRecoilState(serviceListState);
  const [_isLoading, setIsLoading] = useRecoilState(loadingState);

  // Click create button handler.
  const onClickCreate = async () => {
    setIsLoading(true);
    try {
      const result = await fetch(API_SERVICE_CREATE, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          service: form.name,
          artifact: form.artifact,
          options: form.options,
          execute: form.execute,
        }),
      });
      const data = await result.json();
      if (data["message"] === "ok") {
        const response = await fetch(API_GET_SERVICES);
        const data = await response.json();
        setServiceList(data as Service[]);
        setForm(initServiceForm);
      }
    } catch (e) {
      console.error(e);
      alert("Sorry, you got error.");
    } finally {
      setIsLoading(false);
    }
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
          label="Execute command when start contaier."
          variant="outlined"
          value={form.execute}
          onChange={(e) => setForm({ ...form, execute: e.target.value })}
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
