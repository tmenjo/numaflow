import React, { useCallback, useContext, useState } from "react";
import Box from "@mui/material/Box";
import Modal from "@mui/material/Modal";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import { Metrics } from "../../../../pages/Pipeline/partials/Graph/partials/NodeInfo/partials/Pods/partials/PodDetails/partials/Metrics";
import {
  VertexDetailsContext,
  VertexDetailsContextProps,
} from "../../../SlidingSidebar/partials/VertexDetails";
import { Pod } from "../../../../../types/declarations/pods";

import "./style.css";

const modalStyle = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  height: "60%",
  width: "80%",
  backgroundColor: "background.paper",
  boxShadow: 24,
  borderRadius: "1rem",
  p: "3rem",
};

interface MetricsModalProps {
  isModalOpen: boolean;
  handleCloseModal: () => void;
  metricDisplayName: string;
  discoveredMetrics: any;
  namespaceId: string;
  pipelineId: string;
  vertexId: string;
  type: string;
  presets?: any;
  pod?: Pod;
}

export function MetricsModal({
  isModalOpen,
  handleCloseModal,
  metricDisplayName,
  discoveredMetrics,
  namespaceId,
  pipelineId,
  vertexId,
  type,
  presets,
  pod,
}: MetricsModalProps) {
  const { setVertexTab, setPodsViewTab, setExpanded, setPresets } =
    useContext<VertexDetailsContextProps>(VertexDetailsContext);

  const [metricsFound, setMetricsFound] = useState<boolean>(false);

  const handleRedirect = useCallback(() => {
    handleCloseModal();
    if (presets) setPresets(presets);
    setVertexTab(0);
    setPodsViewTab(1);
    // expand the respective metrics accordion
    const discoveredMetric = discoveredMetrics?.data?.find(
      (m: any) => m?.display_name === metricDisplayName
    );
    const panelId = `${discoveredMetric?.metric_name}-panel`;
    setExpanded((prevExpanded) => new Set(prevExpanded).add(panelId));
  }, [
    handleCloseModal,
    presets,
    setPresets,
    setVertexTab,
    setPodsViewTab,
    discoveredMetrics,
    metricDisplayName,
    setExpanded,
  ]);

  return (
    <Modal
      open={isModalOpen}
      onClose={handleCloseModal}
      aria-labelledby="buffer-details-title"
      aria-describedby="buffer-details-description"
    >
      <Box sx={modalStyle}>
        <Box className={"metrics-modal-title-container"}>
          <Box className={"metrics-modal-title"}>{metricDisplayName}</Box>
          <IconButton onClick={handleCloseModal} aria-label="close">
            <CloseIcon fontSize="large" />
          </IconButton>
        </Box>
        <Box sx={{ height: "calc(100% - 3rem)" }}>
          <Metrics
            namespaceId={namespaceId}
            pipelineId={pipelineId}
            vertexId={vertexId}
            type={type}
            metricDisplayName={metricDisplayName}
            setMetricsFound={setMetricsFound}
            presets={presets}
            pod={pod}
          />
        </Box>
        {metricsFound && (
          <Box className={"metrics-modal-redirect"} onClick={handleRedirect}>
            Click to see detailed view with additional filters
          </Box>
        )}
      </Box>
    </Modal>
  );
}
