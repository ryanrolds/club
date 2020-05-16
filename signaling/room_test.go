package signaling_test


var  _ = Describe("Room", func() {
	var (
		fakeConn		*signalingfakes.FakePeerConnection
		Room			*signaling.Room
		validMessage = 	[]byte(`{"type":"type","destId":"destID","payload":{}}`)
	)

	BeforeEach(func() {
		fakeConn = &signalingfakes.FakePeerConnection()
		room = signaling.NewRoom(fakeConn)
	})

	Context("NewRoom", func() {
		It("should create a new room", func() {
			room = signaling.NewRoom(fakeConn)
			Expect(room).ToNot(BeNil())
		})
	})
})