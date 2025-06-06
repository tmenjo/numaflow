// This file is @generated by prost-build.
/// GCEvent is the event that is persisted in the WAL when a window is garbage collected
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GcEvent {
    /// start_time of the window
    #[prost(message, optional, tag = "1")]
    pub start_time: ::core::option::Option<::prost_types::Timestamp>,
    /// end time of the window
    #[prost(message, optional, tag = "2")]
    pub end_time: ::core::option::Option<::prost_types::Timestamp>,
    /// keys of the window, it will be empty for aligned windows
    #[prost(string, repeated, tag = "3")]
    pub keys: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct Footer {
    /// the latest event time in the current Segment
    #[prost(message, optional, tag = "1")]
    pub latest_event_time: ::core::option::Option<::prost_types::Timestamp>,
}
