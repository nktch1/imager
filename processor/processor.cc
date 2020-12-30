#include <iostream>
#include <memory>
#include <sstream>

#include <grpcpp/grpcpp.h>

#include "api.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using api::Processor;

class ProcessorServiceImpl final : public api::Processor::Service {
    Status Process(ServerContext* context, grpc::ServerReaderWriter<api::HumanPosition, api::RawImage>* stream) override {
        api::RawImage req;
        std::ostringstream buf;
        api::HumanPosition hp;

        while (stream->Read(&req)) {
            auto *p = new api::Point;
            p->set_x(atof(req.id().c_str()));
            p->set_y(13);

            std::cout << req.id() << std::endl;
            hp.set_allocated_point(p);
            stream->Write(hp);
        }

        return Status::OK;
    }
};


void Run() {
    std::string address("0.0.0.0:5000");
    ProcessorServiceImpl service;

    ServerBuilder builder;

    builder.AddListeningPort(address, grpc::InsecureServerCredentials());
    builder.RegisterService(&service);

    std::unique_ptr<Server> server(builder.BuildAndStart());
    std::cout << "Server listening on port: " << address << std::endl;

    server->Wait();
}

int main() {
    Run();
}
