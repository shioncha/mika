import { useState } from "react";
import { useParams } from "react-router";

import TagsComponent from "../../elements/tags";
import TimelineComponent from "../../elements/timeline";
import TimelineHeaderComponent from "../../elements/timelineHeader";

function TimelineLayout() {
  const { tag } = useParams();
  const [onlyUnchecked, setOnlyUnchecked] = useState<boolean>(false);
  const [fadeout, setFadeout] = useState<boolean>(false);

  return (
    <main>
      <TimelineHeaderComponent
        onlyUnchecked={onlyUnchecked}
        setOnlyUnchecked={setOnlyUnchecked}
      />
      <TagsComponent tag={tag} setFadeout={setFadeout} />
      <TimelineComponent
        tag={tag}
        fadeout={fadeout}
        setFadeout={setFadeout}
        onlyUnchecked={onlyUnchecked}
      />
    </main>
  );
}

export default TimelineLayout;
