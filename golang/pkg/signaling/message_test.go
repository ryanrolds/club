package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {
	var nodeID = signaling.NodeID("nodeID")
	var validMessage = []byte(`{"type":"type","destId":"destID","payload":{}}`)

	Context("NewJoinMessage", func() {
		It("should return join message", func() {
			message := signaling.NewJoinMessage(signaling.NodeID("123"))
			Expect(message.Type).To(Equal(signaling.MessageTypeJoin))
			Expect(message.SourceID).To(Equal(signaling.NodeID("123")))
		})
	})

	Context("NewJoinMessage", func() {
		It("should return leave message", func() {
			message := signaling.NewLeaveMessage(signaling.NodeID("123"))
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
			Expect(message.SourceID).To(Equal(signaling.NodeID("123")))
		})
	})

	Context("NewMessageFromBytes", func() {
		It("should create message from bytes", func() {
			message, err := signaling.NewMessageFromBytes(nodeID, validMessage)
			Expect(err).To(BeNil())
			Expect(message.Type).To(Equal(signaling.MessageType("type")))
			Expect(message.DestinationID).To(Equal(signaling.NodeID("destID")))
			Expect(message.SourceID).To(Equal(nodeID))
			Expect(message.Payload).To(Equal(signaling.MessagePayload{}))
		})

		It("should error if invalid JSON", func() {
			invalidMessage := []byte(`{`)
			_, err := signaling.NewMessageFromBytes(nodeID, invalidMessage)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("unexpected end of JSON input"))
		})

		It("should error if message is missing type", func() {
			incompleteMessage := []byte(`{"destId":"destID","payload":{}}`)
			_, err := signaling.NewMessageFromBytes(nodeID, incompleteMessage)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(`Key: 'Message.Type' Error:Field validation for 'Type' ` +
				`failed on the 'required' tag`))
		})

		It("should error if message is missing source id", func() {
			incompleteMessage := []byte(`{"type":"type","payload":{}}`)
			_, err := signaling.NewMessageFromBytes(signaling.NodeID(""),
				incompleteMessage)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(`Key: 'Message.SourceID' Error:Field validation for 'SourceID' ` +
				`failed on the 'required' tag`))
		})

		It("should error if message is missing payload", func() {
			incompleteMessage := []byte(`{"type":"type","destId":"destID"}`)
			_, err := signaling.NewMessageFromBytes(nodeID, incompleteMessage)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(`Key: 'Message.Payload' Error:Field validation for 'Payload' ` +
				`failed on the 'required' tag`))
		})
	})

	Context("ToJSON", func() {
		It("should return JSON for the message", func() {
			message, err := signaling.NewMessageFromBytes(nodeID, validMessage)
			Expect(err).To(BeNil())

			data, err := message.ToJSON()
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(`{"type":"type","peerId":"nodeID","destId":"destID","payload":{}}`)))
		})
	})

	Context("GetGroupIDFromMessage", func() {
		It("should return the default group if group key not in payload", func() {
			message := signaling.Message{
				Payload: signaling.MessagePayload{
					signaling.MessagePayloadKeyGroup: "test",
				},
			}

			groupID := signaling.GetGroupIDFromMessage(message, signaling.NodeID("default"))
			Expect(groupID).To(Equal(signaling.NodeID("test")))
		})

		It("should return the group ID if group key present", func() {
			message := signaling.Message{
				Payload: signaling.MessagePayload{},
			}

			groupID := signaling.GetGroupIDFromMessage(message, signaling.NodeID("default"))
			Expect(groupID).To(Equal(signaling.NodeID("default")))
		})
	})
})
