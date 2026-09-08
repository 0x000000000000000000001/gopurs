export const compactEntry = reverse => () => reverse
  ? { label: "alpha", count: 5 }
  : { count: 5, label: "alpha" };

export const compactPayload = () => ({
  values: [-3, 0, 7],
  child: { label: "nested", count: 2 }
});

export const compactScalars = () => ({ number: -1.25, flag: true });

export const compactKeys = record => Object.keys(record).join("|");

export const describeEntryMap = record => `${record.count}:${record.label}`;

export const describePayloadMap = record =>
  `${record.child.count}:${record.child.label}:[${record.values.join(" ")}]`;

export const describeScalarsMap = record => `${record.flag}:${record.number}`;
